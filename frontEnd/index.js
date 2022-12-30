function removeError() {
  let username_style = document.getElementById("username");
  username_style.classList.remove("error");
  let password_style = document.getElementById("password");
  password_style.classList.remove("error");
}
function userlogin() {
  let emailId = document.getElementById("username").value;
  let username_style = document.getElementById("username");
  let emailId_warning = document.getElementById("tempfix");
  let password = document.getElementById("password").value;
  let password_style = document.getElementById("password");
  let returnvalue = document.getElementById("loginButton");

  if (emailId === "") {
    username_style.classList.add("error");
    emailId_warning.style.display = "block";
    emailId_warning.innerHTML = "Email can't be empty";
    setTimeout(disablefunction, 3000);
  }

  if (password === "") {
    password_style.classList.add("error");
    emailId_warning.style.display = "block";
    emailId_warning.innerHTML = "Password can't be empty";
    setTimeout(disablefunction, 3000);
    return;
  }

  returnvalue.addEventListener("click", toCheckValidity(emailId, password));
}

function forgotPassword() {
  const returnvalue = document.getElementById("tempfix");
  returnvalue.style.display = "block";
  returnvalue.innerHTML = "&#9888 Feature under development &#9888";
  setTimeout(disablefunction, 3000);
}
function disablefunction() {
  const returnvalue = document.getElementById("tempfix");
  returnvalue.style.display = "none";
}

async function toCheckValidity(emailId, password) {
  let uuid = await CheckValidity(emailId, password);
  if (uuid != "") {
    window.location.replace("allInstructor.html");
  }
}

async function CheckValidity(username, password) {
  const emailId_warning = document.getElementById("tempfix");
  let response = await fetch(
    `http://localhost:5050/instructor-login/${username}/${password}`
  );
  if (response.status != 500) {
    let uuid_instructor = await response.json();
    return uuid_instructor;
  } else {
    let username_style = document.getElementById("username");
    username_style.classList.add("error");
    let password_style = document.getElementById("password");
    password_style.classList.add("error");
    emailId_warning.style.display = "block";
    emailId_warning.innerHTML = "Wrong Credentials";
    setTimeout(disablefunction, 3000);
    return "";
  }
}
