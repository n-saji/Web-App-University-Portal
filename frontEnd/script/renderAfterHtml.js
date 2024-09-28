const d = new Date();
let time = d.getHours();
if (time > 0 && time < 12) {
  let h1 = document.querySelector(".greeting-message");
  h1.innerHTML = "Good Morning!";
} else if (time >= 12 && time <= 18) {
  let h1 = document.querySelector(".greeting-message");
  h1.innerHTML = "Good AfterNoon!";
} else {
  let h1 = document.querySelector(".greeting-message");
  h1.innerHTML = "Good Evening!";
}
