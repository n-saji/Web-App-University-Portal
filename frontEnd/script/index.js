// check if cookie present , validate it
let cookies = document.cookie.split(";");
let cookiesMap = new Map();
for (let index = 0; index < cookies.length; index++) {
  let cookie = cookies[index].split("=");
  cookie[0] = cookie[0].replace(" ", "");
  cookiesMap[cookie[0]] = cookie[1];
}

setTimeout(validateCookie, 1000);

async function validateCookie() {
  let response = await fetch(`http://localhost:5050/check-token-status`, {
    method: "GET",
    headers: {
      Token: cookiesMap["token"],
    },
  });
  console.log(response);
  let jsonResponse = await response.json();
  console.log(jsonResponse);
  if (response.status != 500) {
    window.location.replace("dashboard-v2.html");
  }
}


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

  if (!emailId) {
    //username_style.style.border = "2px solid red";
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
    window.location.replace("dashboard-v2.html");
  }
}

async function CheckValidity(username, password) {
  const emailId_warning = document.getElementById("tempfix");
  let error_while_fetching_api;
  let response = await fetch(`http://localhost:5050/v1/login`, {
    method: "POST",
    body: JSON.stringify({
      email_id: username,
      password: password,
    }),
  }).catch((err) => {
    error_while_fetching_api = err;
  });
  if (error_while_fetching_api == "TypeError: Failed to fetch") {
    let error = document.getElementById("tempfix");
    error.style.display = "block";
    error.innerHTML = "server down &#9760;";
    setTimeout(disablefunction, 4000);
    return "";
  }
  if (response.status != 500) {
    let uuid_instructor = await response.json();
    document.cookie =
      "token" + "=" + response.headers.get("token") + "; path=/";
    document.cookie =
      "account_id" + "=" + response.headers.get("account_id") + "; path=/";
    return uuid_instructor;
  } else {
    let response_reply = await response.json();
    console.log(response_reply);
    let username_style = document.getElementById("username");
    username_style.classList.add("error");
    let password_style = document.getElementById("password");
    password_style.classList.add("error");
    emailId_warning.style.display = "block";
    emailId_warning.innerHTML = response_reply;
    setTimeout(disablefunction, 3000);
    return "";
  }
}

function showPassword() {
  var img_src = document.getElementById("image_for_show_password");
  var password_img = document.getElementById("password_img");
  var x = document.getElementById("password");
  var hide_img = document.getElementById("hide_img");

  if (x.type === "password") {
    x.type = "text";
    var img_val = document.createElement("img");
    img_val.src = "/frontEnd/asset/view.png";
    img_val.classList.add("image_for_show_password");
    img_val.onclick = showPassword;
    img_val.id = "hide_img";
    img_src.classList.add("hide_element");
    password_img.appendChild(img_val);
  } else {
    x.type = "password";
    if (hide_img != null) {
      hide_img.remove();
    }
    img_src.classList.remove("hide_element");
  }
}

function signup() {
  window.location.replace("signup.html");
}
