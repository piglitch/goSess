// JavaScript to change the background color
const Body = document.body;
Body.style.backgroundColor = 'beige';

const newButton = document.createElement('button');
newButton.innerText = "Click to change the background"

newButton.addEventListener("click", () => {
  if (Body.style.backgroundColor === 'skyblue') {
    Body.style.backgroundColor = 'beige';
    return;
  }
  Body.style.backgroundColor = 'skyblue' 
})

Body.appendChild(newButton);

const newButton2 = document.createElement('button');
newButton2.innerText = "Click to send a post request"

Body.appendChild(newButton2);

newButton2.addEventListener('click', () => {
  fetchReq();
});


const fetchReq = async () => {
  const data = {
    name: "John Doe",
    email: "john@example.com"
  };
  try {
    const response = await fetch('http://localhost:8080/api/json', {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json'
      },
      body: JSON.stringify(data)
    })
    if (response.ok) {
      const jsonResponse = await response.json();
      console.log(jsonResponse);
      console.log(data);
    }
  } catch (error) {
    alert("err: ", err)
  }
}