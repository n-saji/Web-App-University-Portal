function userlogin() {
  const emailId = document.getElementById("username").value;
  const username_style = document.getElementById("username");
  const emailId_warning = document.getElementById("tempfix");

  if (!emailId) {
    username_style.style.border = "2px solid red";
    emailId.innerHTML = "!";
    emailId_warning.style.display = "block";
    emailId_warning.innerHTML = "Email can't be empty";
    setTimeout(disablefunction, 3000);
    return;
  }
  const password = document.getElementById("password").value;
  const password_style = document.getElementById("password");
  if (!password) {
    password_style.style.border = "2px solid red";
    emailId_warning.style.display = "block";
    emailId_warning.innerHTML = "Password can't be empty";
    setTimeout(disablefunction, 3000);
    return;
  }
  const returnvalue = document.getElementById("loginButton");

  returnvalue.addEventListener("click", toCheckValidity(emailId, password));
}

function forgotPassword() {
  const returnvalue = document.getElementById("tempfix");
  returnvalue.style.display = "block";
  returnvalue.innerHTML = "&#9888Feature under development&#9888";
  setTimeout(disablefunction, 3000);
}
function disablefunction() {
  const returnvalue = document.getElementById("tempfix");
  returnvalue.style.display = "none";
}

async function toCheckValidity(emailId, password) {
  let uuid = await CheckValidity(emailId, password).catch((error) =>
    console.log(error)
  );
  if (uuid != "") {
    console.log("correct", uuid);
    window.location.replace("allInstructor.html");
  }
}

async function CheckValidity(username, password) {
  const emailId_warning = document.getElementById("tempfix");
  let response = await fetch(
    `http://localhost:5050/instructor-login/${username}/${password}`
  );
  let uuid_instructor = await response.json();
  if (response.ok == true) {
    return uuid_instructor;
  } else {
    emailId_warning.style.display = "block";
    emailId_warning.innerHTML = "Wrong Credentials";
    setTimeout(disablefunction, 3000);
    return "";
  }
}
