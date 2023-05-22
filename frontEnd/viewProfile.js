async function getInstructorDetails() {
  let cookie_token = getCookie("token");
  let instructor_id = getCookie("account_id");
  let api_error;
  let getDetails = await fetch(
    `http://localhost:5050/get-instructor-name-by-id/${instructor_id}`,
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
  if (response == "token expired! Generate new token") {
    alert("Timed-out re login");
    setTimeout(window.location.replace("index.html"), 2000);
    return;
  }
  let i_name = document.getElementById("instructor_name");

  i_name.innerHTML = response.instructor_name;
  autofil(instructor_id);
}
function getCookie(name) {
  var nameEQ = name + "=";
  var ca = document.cookie.split(";");
  for (var i = 0; i < ca.length; i++) {
    var c = ca[i];
    while (c.charAt(0) == " ") c = c.substring(1, c.length);
    if (c.indexOf(nameEQ) == 0) return c.substring(nameEQ.length, c.length);
  }
  return null;
}
getInstructorDetails();
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
function setdashboard() {
  window.location.replace("dashboard.html");
}

async function autofil(id_instructor) {
  let api_url = "";
  let url = `http://localhost:5050/view-profile-instructor/${id_instructor}`;
  let cookie_token = getCookie("token");
  let api_response = await fetch(url, {
    method: "GET",
    headers: { Token: cookie_token },
  }).catch((err) => {
    api_url = err;
  });
  if (api_url != "") {
    alert("Internal Server Error Please Login Again" + api_url);
    return "";
  }
  let response = await api_response.json();
  if (response == "token expired! Generate new token") {
    alert("Timed-out re login");
    window.location.replace("index.html");
    return;
  }
  console.log(response);
  let name_input = document.getElementById("name_input");
  name_input.value = response.Name;
  let email_input = document.getElementById("email_input");
  email_input.value = response.Credentials.email_id;
  let password_input = document.getElementById("password_input");
  password_input.value = response.Credentials.password.slice(0, 9);
  let code_input = document.getElementById("code_input");
  code_input.value = response.Code;
}
