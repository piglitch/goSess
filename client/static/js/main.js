// JavaScript to change the background color
const Body = document.body;
Body.style.backgroundColor = 'beige';

const newButton = document.createElement('button');
newButton.innerText = "Click to change the background"

newButton.addEventListener("click", () => {
  if (Body.style.backgroundColor === 'teal') {
    Body.style.backgroundColor = 'beige';
    return;
  }
  Body.style.backgroundColor = 'teal' 
})

Body.appendChild(newButton);
