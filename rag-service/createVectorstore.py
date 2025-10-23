import pandas as pd
from langchain_text_splitters import RecursiveCharacterTextSplitter
from langchain_community.vectorstores import FAISS
from langchain_huggingface import HuggingFaceEmbeddings

from dotenv import load_dotenv
from sentence.functions import make_exercise_sentence, make_food_sentence

load_dotenv()

class DataFiles():
    def __init__(self, file_paths=[], tag="", sentence_fn=None):
        self.file_paths = file_paths
        self.tag = tag
        self.sentence_fn = sentence_fn
        self.normalized_data_path = ""

def combine_data(file_paths):
    dfs = [pd.read_csv(fp) for fp in file_paths]
    df = pd.concat(dfs, ignore_index=True)

    return df

def normalize_data(file):
    if not file.file_paths:
        raise Exception("No paths provided")
    
    # Load the Kaggle dataset
    if len(file.file_paths) > 1:
        df = combine_data(file.file_paths)
    else:
        df = pd.read_csv(file.file_paths[0])

    print(f"Combined dataset size: {len(df)} rows")

    # Apply to every row
    df["text"] = df.apply(file.sentence_fn, axis=1)

    # Save text data for RAG ingestion
    file.normalized_data_path = f"data/{file.tag}/{file.tag}_text.csv"

    df[["text"]].to_csv(file.normalized_data_path, index=False)

def embedd_data(file):
    normalize_data(file)

    df = pd.read_csv(file.normalized_data_path)
    texts = df["text"].tolist()

    splitter = RecursiveCharacterTextSplitter(chunk_size=1000, chunk_overlap=100)
    docs = splitter.create_documents(texts)

    embeddings = HuggingFaceEmbeddings(model_name="sentence-transformers/all-MiniLM-L6-v2")
    db = FAISS.from_documents(docs, embeddings)
    db.save_local(f"vectors/{file.tag}")

if __name__ == "__main__":
    files = []
    files.append(
        DataFiles(
            file_paths = [
                "./data/food/FINAL_FOOD_DATASET/FOOD-DATA-GROUP1.csv",
                "./data/food/FINAL_FOOD_DATASET/FOOD-DATA-GROUP2.csv",
                "./data/food/FINAL_FOOD_DATASET/FOOD-DATA-GROUP3.csv",
                "./data/food/FINAL_FOOD_DATASET/FOOD-DATA-GROUP4.csv",
                "./data/food/FINAL_FOOD_DATASET/FOOD-DATA-GROUP5.csv",
            ],
            tag = "food",
            sentence_fn = make_food_sentence
        )
    )

    files.append(
        DataFiles(
            file_paths = ["./data/activity/exercise_dataset.csv"],
            tag = "activity",
            sentence_fn = make_exercise_sentence
        )
    )
    for file in files:
        embedd_data(file)
