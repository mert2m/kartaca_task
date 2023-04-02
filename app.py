from flask import Flask, jsonify
from pymongo import MongoClient

app = Flask(__name__)

MONGO_URI = 'mongodb://mongodb:27017/'
client = MongoClient(MONGO_URI)
db = client.stajdb
iller = db.iller

@app.route("/")
def hello():
    return "Merhaba Python!"

@app.route("/staj")
def get_random_il():
    il = iller.aggregate([{ '$sample': { 'size': 1 } }]).next()
    il.pop('_id')
    return jsonify(il)

if __name__ == "__main__":
    iller.insert_one({"ad": "Adana"})
    iller.insert_one({"ad": "Ankara"})
    iller.insert_one({"ad": "İstanbul"})
    iller.insert_one({"ad": "İzmir"})
    iller.insert_one({"ad": "Düzce"})
    iller.insert_one({"ad": "Sakarya"})
    iller.insert_one({"ad": "Tekirdağ"})
    iller.insert_one({"ad": "Edirne"})
    iller.insert_one({"ad": "Antalya"})
    iller.insert_one({"ad": "Gaziantep"})
    iller.insert_one({"ad": "Trabzon"})
    app.run(debug=True, host='0.0.0.0', port=4444)
