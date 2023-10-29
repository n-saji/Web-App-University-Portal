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
  window.location.replace("dashboard-v2.html");
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
  let name_input = document.getElementById("name_input");
  name_input.value = response.Name;
  let email_input = document.getElementById("email_input");
  email_input.value = response.Credentials.email_id;
  let password_input = document.getElementById("password_input");
  password_input.value = response.Credentials.password.slice(0, 9);
  let code_input = document.getElementById("code_input");
  code_input.value = response.Code;
}

async function ToUpdateDetails(id_name, type) {
  let instructor_id = getCookie("account_id");
  let cookie_token = getCookie("token");

  if (type == "name") {
    let req_name_id = document.getElementById(id_name);
    let req_name = req_name_id.value;

    let api_url = "";
    let url = `http://localhost:5050/update-instructor?instructor_id=${instructor_id}`;
    let api_response = await fetch(url, {
      method: "PATCH",
      headers: { Token: cookie_token },
      body: JSON.stringify({
        instructor_name: req_name,
      }),
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
    if (api_response.status != 500) {
      window.location.reload();
    }
  }
  if (type == "code") {
    let req_id = document.getElementById(id_name);
    let req_value = req_id.value;
    let api_url = "";
    let url = `http://localhost:5050/update-instructor?instructor_id=${instructor_id}`;
    let api_response = await fetch(url, {
      method: "PATCH",
      headers: { Token: cookie_token },
      body: JSON.stringify({
        instructor_code: req_value,
      }),
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
    if (api_response.status != 500) {
      window.location.reload();
    }
  }
  if (type == "email") {
    let req_id = document.getElementById(id_name);
    let req_value = req_id.value;
    let api_url = "";
    let url = `http://localhost:5050/update-instructor-credentials`;
    let api_response = await fetch(url, {
      method: "PUT",
      headers: { Token: cookie_token },
      body: JSON.stringify({
        id: instructor_id,
        email_id: req_value,
      }),
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
    if (api_response.status != 500) {
      window.location.reload();
    }
  }
  if (type == "password") {
    let req_id = document.getElementById(id_name);
    let req_value = req_id.value;
    let api_url = "";
    let url = `http://localhost:5050/update-instructor-credentials`;
    let api_response = await fetch(url, {
      method: "PUT",
      headers: { Token: cookie_token },
      body: JSON.stringify({
        id: instructor_id,
        password: req_value,
      }),
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
    if (api_response.status != 500) {
      window.location.reload();
    }
  }
}

function insertInstructor() {
  window.location.replace("createInstructor.html");
}

function logout() {
  fetch(`http://localhost:5050/logout?token=${getCookie("token")}`, {});
  window.location.replace("index.html");
}
function setViewProfile() {
  window.location.replace("viewProfile.html");
}

function insertCourse() {
  window.location.replace("createCourse.html");
}

function showCourse() {
  window.location.replace("showCourse.html");
}

function insertStudent() {
  window.location.replace("addStudent.html");
}

function showCourse() {
  window.location.replace("showCourse.html");
}

function showStudent() {
  window.location.replace("showStudents.html");
}

function showInstructor() {
  window.location.replace("instructorDetails.html");
}

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
  let i_name = document.querySelector(".logged-user");
  i_name.innerHTML = response.instructor_name;
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
