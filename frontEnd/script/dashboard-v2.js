function insertInstructor() {
  window.location.replace("createInstructor.html");
}

function setdashboard() {
  window.location.replace("dashboard-v2.html");
}
function logout() {
  fetch(`http://3.111.149.112:5050/logout?token=${getCookie("token")}`, {});
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
    `http://3.111.149.112:5050/get-instructor-name-by-id/${instructor_id}`,
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
  let i_name = document.querySelector(".logged-user");
  i_name.innerHTML = response.instructor_name;
}

getInstructorDetails();
