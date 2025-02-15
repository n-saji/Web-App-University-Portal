function showPassword() {
  var x = document.querySelector(".input-password");

  if (x.type === "password") {
    x.type = "text";
  } else {
    x.type = "password";
  }
}
function signup() {
  window.location.replace("signup.html");
}

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
  let response = await fetch(
    `https://dolphin-app-2zya2.ondigitalocean.app/check-token-status`,
    {
      method: "GET",
      headers: {
        Token: cookiesMap["token"],
      },
    }
  );
  console.log(response);
  let jsonResponse = await response.json();
  console.log(jsonResponse);
  if (response.status != 500) {
    window.location.replace("dashboard-v2.html");
  }
}

function userlogin() {
  let emailId = document.querySelector(".input-email");
  let password = document.querySelector(".input-password");
  if (emailId.value === "") {
    let errorPara = document.querySelector(".error-message");
    errorPara.innerHTML = "Empty Email!";
    let errorPopup = document.querySelector(".error-popup");
    errorPopup.classList.add("error-popup-display");
    return;
  }

  if (password.value === "") {
    let errorPara = document.querySelector(".error-message");
    errorPara.innerHTML = "Empty Password!";
    let errorPopup = document.querySelector(".error-popup");
    errorPopup.classList.add("error-popup-display");
    return;
  }

  toCheckValidity(emailId.value, password.value);
}

async function toCheckValidity(emailId, password) {
  let uuid = await CheckValidity(emailId, password);
  if (uuid != "") {
    await getInstructorDetails();
    console.log(localStorage.getItem("username"));
    window.location.replace("dashboard-v2.html");
  }
}

async function CheckValidity(username, password) {
  let error_while_fetching_api;
  let response = await fetch(
    `https://dolphin-app-2zya2.ondigitalocean.app/v1/login`,
    {
      method: "POST",
      body: JSON.stringify({
        email_id: username,
        password: password,
      }),
    }
  ).catch((err) => {
    error_while_fetching_api = err;
  });

  if (error_while_fetching_api == "TypeError: Failed to fetch") {
    let errorPara = document.querySelector(".error-message");
    errorPara.innerHTML = "Server down &#9760 ";
    let errorPopup = document.querySelector(".error-popup");
    errorPopup.classList.add("error-popup-display");
    setTimeout(removeError, 4000);
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
    let errorPara = document.querySelector(".error-message");
    errorPara.innerHTML = "Wrong Credentials!";
    let errorPopup = document.querySelector(".error-popup");
    errorPopup.classList.add("error-popup-display");
    // setTimeout(removeError, 5000);
    return "";
  }
}

function removeError() {
  let errorPopup = document.querySelector(".error-popup");
  errorPopup.classList.remove("error-popup-display");
}

addEventListener("keypress", removeError);
addEventListener("keypress", function (e) {
  console.log(e);
  if (e.key == "Enter") {
    console.log(e.key);
    userlogin();
  }
});

async function getInstructorDetails() {
  let cookie_token = getCookie("token");
  let instructor_id = getCookie("account_id");
  let api_error;
  let getDetails = await fetch(
    `https://dolphin-app-2zya2.ondigitalocean.app/get-instructor-name-by-id/${instructor_id}`,
    {
      method: "GET",
      headers: { Token: cookie_token },
    }
  ).catch((err) => {
    api_error = err;
  });
  if (api_error == "TypeError: Failed to fetch") {
    alert("Internal Server Error Please Login Again");
    window.location.replace("index.html");
    return "";
  }
  let response = await getDetails.json();
  localStorage.setItem("username", response.instructor_name);
}

function reloadToSignUp() {
  window.location.replace("signup.html");
}
