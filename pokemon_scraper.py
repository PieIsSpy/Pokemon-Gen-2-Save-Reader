import requests
import pandas as pd

url_stats = "https://bulbapedia.bulbagarden.net/wiki/List_of_Pok%C3%A9mon_by_base_stats_in_Generations_II-V"
url_index = "https://bulbapedia.bulbagarden.net/wiki/List_of_Pok%C3%A9mon_by_index_number_in_Generation_II"

header = {
  "User-Agent": "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/50.0.2661.75 Safari/537.36",
  "X-Requested-With": "XMLHttpRequest"
}

# request for the pages
req1 = requests.get(url_stats, headers=header)
req2 = requests.get(url_index, headers=header)

# get the table for pokemon gen 2 indices
df_index = pd.read_html(req2.text)[0]
df_index.columns = df_index.columns.get_level_values(1)
df_index = df_index.drop(columns=["HEX", "MS"])
df_index = df_index.rename(columns={"Name": "Pokemon"})
df_index['Pokemon'] = df_index['Pokemon'].str.replace('Nidoran♂', 'NidoranM')
df_index['Pokemon'] = df_index['Pokemon'].str.replace('Nidoran♀', 'NidoranF')

# get the table for pokemon base stats
df_stats = pd.read_html(req1.text)[1]
df_stats['Pokémon'] = df_stats["Pokémon.1"]
df_stats = df_stats.rename(columns={"Pokémon": "Pokemon"})
df_stats = df_stats.drop(columns=["Pokémon.1", "Total", "Average"])
df_stats['Pokemon'] = df_stats['Pokemon'].str.replace('Nidoran♂', 'NidoranM')
df_stats['Pokemon'] = df_stats['Pokemon'].str.replace('Nidoran♀', 'NidoranF')

# combine all tables
df = pd.merge(df_stats, df_index, on='Pokemon')
df = df.rename(columns={"#": "DexNum", "Types": "Type1", "Types.1": "Type2"})
df = df.loc[:, ["DEC", "DexNum", "Pokemon", "Type1", "Type2", "HP",  "Attack",  "Defense",  "Speed",  "Sp. Atk", "Sp. Def"]]
df = df.set_index("DEC")
df = df.sort_values(by=["DEC"])

df.to_csv('pkmn2.csv')