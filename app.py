from selenium import webdriver
from selenium.webdriver.chrome.service import Service
from webdriver_manager.chrome import ChromeDriverManager as cm
from selenium.webdriver.common.by import By
from dotenv import load_dotenv
import time
import os

load_dotenv()

URL = os.getenv("URL")
XPATH = os.getenv("XPATH")

service = Service(executable_path=cm().install())

driver = webdriver.Chrome(service=service)
driver.get(URL)
time.sleep(5)

for i in range(10):
    driver.execute_script("window.scrollTo(0, 2000);")
    time.sleep(0.5)

tweets = driver.find_elements(By.XPATH, XPATH)

for tweet in tweets:
    print(tweet.text)
