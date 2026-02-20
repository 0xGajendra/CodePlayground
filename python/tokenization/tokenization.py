from tiktoken import encoding_for_model

enc = encoding_for_model("gpt-4o")
# text = "Hey there! My name is gajendra"

# token = enc.encode(text)

print(enc.decode([25216, 1354, 0, 3673, 1308, 382, 329, 1255, 32364]))