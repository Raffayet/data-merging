import easyocr
import re
import spacy

# Load spaCy's pre-trained NER model
nlp = spacy.load("en_core_web_sm")

# Initialize the EasyOCR reader
reader = easyocr.Reader(['en'])  # Specify the language(s) for OCR

# Function to extract email from text using regex
def extract_email(text):
    # Regex for email detection, allowing spaces around email parts
    email_pattern = r'[a-zA-Z0-9._%+-]+\s*@\s*[a-zA-Z0-9.-]+\s*\.\s*[a-zA-Z]{2,}'
    
    # Search for the email pattern in the text
    email_match = re.search(email_pattern, text)
    
    if email_match:
        # Extract the matched email and remove spaces **only** within the email
        raw_email = email_match.group(0)
        cleaned_email = re.sub(r'\s+', '', raw_email)  # Remove spaces from the detected email
        return cleaned_email
    
    return None

# Function to extract phone number from text using regex
def extract_phone_number(text):
    phone_pattern = r'\b\d{9,14}\b'
    phone_match = re.search(phone_pattern, text)
    if phone_match:
        return phone_match.group(0)
    return None

# Function to extract address from text using regex and fallback to spaCy NER
def split_before_address(text):
    # Remove the parts before the first known address boundary (e.g., after the phone number or email)
    # You can split based on known phone or email patterns
    
    # Example: remove everything before the phone number
    text_parts = re.split(r'\b\d{9,14}\b', text)  # Splits at the phone number
    if len(text_parts) > 1:
        return text_parts[1]  # Return the part after the phone number
    return text  # Return original text if no split happened

def extract_address(text):
    # Call split_before_address to clean the text before extracting address
    cleaned_text = split_before_address(text)

    # Use the same regex as before
    address_pattern = r'(\b(?:\d{1,5}\s)?(?:[A-Z][a-z]+\s)+(Street|St|Avenue|Ave|Boulevard|Blvd|Road|Rd|Lane|Ln|Drive|Dr|Way|Court|Ct|Place|Pl),\s*\w+(?:\s\w+)*,\s*[A-Z]{2,3}\s*\d{5}(?:-\d{4})?,?\s*(?:USA|United States|UK|Canada)?)'
    
    # Try extracting address via regex
    address_match = re.search(address_pattern, cleaned_text)
    if address_match:
        return address_match.group(0)
    
    # Fallback: use spaCy NER if regex fails
    doc = nlp(cleaned_text)
    for ent in doc.ents:
        if ent.label_ in ("GPE", "LOC", "FAC", "ORG"):
            return ent.text

    return None


# Function to extract names using spaCy NER
def extract_names(text):
    doc = nlp(text)
    first_name = None
    # Use spaCy to detect the first name only
    for ent in doc.ents:
        if ent.label_ == "PERSON":
            first_name = ent.text
            break  # Only get the first name
    return first_name

# Function to split full name into first name and last name
def split_first_last_name(full_name):
    parts = full_name.split()
    first_name = parts[0]  # The first part is the first name
    last_name = " ".join(parts[1:]) if len(parts) > 1 else ""  # Join the remaining parts as the last name, if any
    return first_name, last_name

# Function to process the image, extract text, and categorize details
def process_image_and_categorize(photo_path):
    # Read the text from the image using EasyOCR
    result = reader.readtext(photo_path)
    extracted_text = " ".join([detection[1] for detection in result])

    print("Full Extracted Text:", extracted_text)

    # Extracting details
    email = extract_email(extracted_text)
    phone_number = extract_phone_number(extracted_text)
    name = extract_names(extracted_text)
    first_name, last_name = split_first_last_name(name)
    address = extract_address(extracted_text)

    # Display the results
    print(f"First Name: {first_name}")
    print(f"Last Name: {last_name}")
    print(f"Email: {email}")
    print(f"Phone Number: {phone_number}")
    print(f"Address: {address}")

# Main block to run the script and process the image
if __name__ == "__main__":
    photo_path = 'data/photo_input.png'  # Path to your image
    process_image_and_categorize(photo_path)
