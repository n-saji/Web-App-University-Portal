async function createEmail() {
  var url = localStorage.getItem("URL_Create_Login");
  let url_sliced = url.slice(0, 84);

  let emailid = document.getElementById("emailid").value;
  if (!emailid) {
    alert("Please enter emailid.");
    return;
  }
  let password = document.getElementById("password").value;
  if (!password) {
    alert("Please enter password.");
    return;
  }
  let url_final = url_sliced + emailid + "/" + password;
  let response = await fetch(url_final);
  let response_reply = await response.json();
  let reply_for_login = document.getElementById("response_for_login");
  if (!response.ok) {
    reply_for_login.innerHTML = response_reply + "!!";
  } else {
    reply_for_login.innerHTML = "Successffully Created";
  }
}
function setdashboard() {
  window.location.replace("dashboard.html");
}
function setbackpage() {
  window.location.replace("index.html");
}

setInterval(checkTokenValidity, 300000);

async function checkTokenValidity() {
  let api_error;
  let cookie_token = getCookie("token");
  let api_response = await fetch(`http://localhost:5050/check-token-status`, {
    method: "GET",
    headers: { Token: cookie_token },
  }).catch((err) => {
    api_error = err;
  });
  if (api_error == "TypeError: Failed to fetch") {
    alert("Internal Server Error Please Login Again");
    window.location.replace("index.html");
    return "";
  }
  let response = await api_response.json();
  if (response == "token expired! Generate new token") {
    alert("Timed-out re login");
    window.location.replace("index.html");
    return;
  }
}