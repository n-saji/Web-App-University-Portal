function populateInstructors() {
  window.location.replace("instructorDetails.html");
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

function setdashboard() {
  window.location.replace("dashboard-v2.html");
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
  let i_name = document.getElementById("instructor_name");

  i_name.innerHTML = response.instructor_name;
}

getInstructorDetails();
