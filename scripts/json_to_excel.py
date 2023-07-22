import pandas as pd

def json_excel(json_file, excel_file):
    # Read the JSON file
    df = pd.read_json(json_file)

    # Convert the DataFrame to Excel
    df.to_excel(excel_file, index=False)

# Example usage
json_file_path = '../output.json'
excel_file_path = '../output.xlsx'

json_excel(json_file_path, excel_file_path)
