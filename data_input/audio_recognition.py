import re
import pyaudio
import wave
import speech_recognition as sr

# Audio recording settings
FORMAT = pyaudio.paInt16
CHANNELS = 1
RATE = 16000
CHUNK = 1024
RECORD_SECONDS = 15
WAVE_OUTPUT_FILENAME = "data/output.wav"

# Initialize the recognizer for speech recognition
recognizer = sr.Recognizer()

def record_audio():
    audio = pyaudio.PyAudio()

    # Start recording
    stream = audio.open(format=FORMAT, channels=CHANNELS,
                        rate=RATE, input=True,
                        frames_per_buffer=CHUNK)

    print("Snimam zvuk...")

    frames = []
    for i in range(0, int(RATE / CHUNK * RECORD_SECONDS)):
        data = stream.read(CHUNK)
        frames.append(data)

    print("Snimanje završeno.")

    # Stop recording
    stream.stop_stream()
    stream.close()
    audio.terminate()

    # Save audio to a WAV file
    waveFile = wave.open(WAVE_OUTPUT_FILENAME, 'wb')
    waveFile.setnchannels(CHANNELS)
    waveFile.setsampwidth(audio.get_sample_size(FORMAT))
    waveFile.setframerate(RATE)
    waveFile.writeframes(b''.join(frames))
    waveFile.close()

def audio_to_text(filename):
    # Load the audio file and convert it to text
    with sr.AudioFile(filename) as source:
        audio_data = recognizer.record(source)
        try:
            text = recognizer.recognize_google(audio_data)
            print("Recognized text:", text)
            return text
        except sr.UnknownValueError:
            print("Google Speech Recognition could not understand audio")
        except sr.RequestError as e:
            print("Could not request results from Google Speech Recognition service; {0}".format(e))
    return None

def extract_first_name(text):
    # Look for the keyword "first name" and extract the word that follows it
    words = text.split()
    if "first" in words and "name" in words:
        first_name_index = words.index("name") + 1  # Get the word after "name"
        if first_name_index < len(words):
            return words[first_name_index]  # The next word is the name
    return None

def extract_last_name(text):
    # Look for the keyword "last name" and extract the word that follows it
    words = text.split()
    if "last" in words and "name" in words:
        last_name_index = words.index("name") + 1  # Get the word after "name"
        if last_name_index < len(words):
            return words[last_name_index]  # The next word is the name
    return None

def extract_email(text):
    # Use regex to find email addresses in the text
    email_pattern = r'[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}'
    email_match = re.search(email_pattern, text)
    if email_match:
        return email_match.group(0)  # Return the matched email
    return None

if __name__ == "__main__":
    # Step 1: Record the audio
    record_audio()

    # Step 2: Convert the audio to text
    text = audio_to_text(WAVE_OUTPUT_FILENAME)

    if text:
        # Step 3: Extract the first name, last name, and email from the recognized text
        first_name = extract_first_name(text)
        last_name = extract_last_name(text)
        email = extract_email(text)

        if first_name:
            print("Extracted first name:", first_name)
        else:
            print("No first name found in the text.")
        if last_name:
            print("Extracted last name:", last_name)
        else:
            print("No last name found in the text.")

        if email:
            print("Extracted email:", email)
        else:
            print("No email found in the text.")
