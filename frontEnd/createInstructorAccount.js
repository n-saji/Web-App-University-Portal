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
  window.location.replace("allinstructor.html");
}
function setbackpage() {
  window.location.replace("index.html");
}
