const targetNumber = Math.floor(Math.random() * 1000) + 1;
let attempts = 0;
var hi=0, lo=1000;
function checkGuess() {
  const guessInput = document.getElementById('guessInput');
  const guess = parseInt(guessInput.value);

  if (isNaN(guess) || guess < 1 || guess > 1000) {
    alert("Please enter a valid number between 1 and 100.");
    return;
  }

  attempts++;

  if (guess === targetNumber) {
    showResult(`Congratulations! You guessed the correct number ${targetNumber} in ${attempts} attempts.`);
    disableInputAndButton();
  } else if (guess < targetNumber) {
    hi = guess;
    showResult(`${hi} - ${lo}`);
  } else {
    lo=guess;
    showResult(`${hi} - ${lo}`);
  }
  guessInput.value = '';
}

function showResult(message) {
  const result = document.getElementById('result');
  result.textContent = message;
}

function disableInputAndButton() {
  const guessInput = document.getElementById('guessInput');
  const checkButton = document.querySelector('button');
  guessInput.disabled = true;
  checkButton.disabled = true;
}