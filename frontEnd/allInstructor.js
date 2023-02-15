function populateInstructors() {
  window.location.replace("instructorDetails.html");
}

function insertInstructor() {
  window.location.replace("createInstructor.html");
}

function setbackpage() {
  window.location.replace("index.html");
}

function setdashboard() {
  window.location.replace("allinstructor.html");
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

function showStudents() {
  window.location.replace("showStudents.html");
}

async function getInstructorDetails() {
  let cookie_token = getCookie("token");
  let instructor_id = getCookie("account_id");
  let getDetails = await fetch(
    `http://localhost:5050/get-instructor-name-by-id/${instructor_id}`,
    {
      method: "GET",
      headers: { Token: cookie_token },
    }
  );
  let response = await getDetails.json();
  if (response == "token expired! Generate new token") {
    alert("Timed-out re login");
    setTimeout(window.location.replace("index.html"), 2000);
    return;
  }
  let i_name = document.getElementById("instructor_name");

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
