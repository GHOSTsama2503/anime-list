#!/usr/bin/env python3
import argparse
import os
import sqlite3


def read_schema(path: str) -> str:
    with open(path, "rb") as file:
        content = file.read()
        return content.decode(encoding="utf-8")


def delete_db(path: str):
    if os.path.exists(path=path):
        os.remove(path=path)


def main(schema: str, outout: str):
    con = sqlite3.connect(outout)
    sql = read_schema(path=schema)
    con.executescript(sql)
    con.close()


if __name__ == "__main__":
    parser = argparse.ArgumentParser("anime-list", description="generate database from schema")

    parser.add_argument("--schema", "-s", help="path to schema.sql")
    parser.add_argument("--output", "-o", help="where do you want to store the .db")
    parser.add_argument("--delete", "-d", action="store_true", help="delete and create the database again")

    args = parser.parse_args()

    if args.delete:
        delete_db(path=args.output)

    main(schema=args.schema, outout=args.output)
