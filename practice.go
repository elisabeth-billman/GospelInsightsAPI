import axios from 'axios';

const apiUrl = 'http://localhost:8000/ask';

const payload = {
  model: 'gpt-3.5-turbo',
  messages: [
    { role: 'system', content: 'You are a helpful assistant.' },
    { role: 'user', content: 'What is the capital of France?' }
  ]
};

axios.post(apiUrl, payload)
  .then(response => {
    console.log(response.data);
  })
  .catch(error => {
    console.error(error);
  });



  curl -X POST -H "Content-Type: application/json" -d '{
  "model": "gpt-3.5-turbo",
  "messages": [
    {"role": "system", "content": "You are a helpful assistant."},
    {"role": "user", "content": "What is the capital of France?"}
  ]
}' http://localhost:8000/ask


curl -X POST -H "Content-Type: application/json" -d '{
  "model": "gpt-3.5-turbo",
  "messages": [
    {"role": "system", "content": "You are a helpful assistant."},
    {"role": "user", "content": "What is the book of mormon?"}
  ]
}' http://localhost:8000/ask