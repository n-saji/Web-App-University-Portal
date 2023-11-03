function setdashboard() {
  window.location.replace("dashboard-v2.html");
}
function setbackpage() {
  window.location.replace("showStudents.html");
}

async function InsertStudentValues() {
  let student_name = document.getElementById("name");
  let student_age = document.getElementById("age");
  let student_roll_number = document.getElementById("roll_number");
  let student_course = document.getElementById("cn_drop_down");
  let error_log = document.getElementById("error_log");
  var submitStudent = document.getElementById("submitStudent");
  if (student_name.value === "") {
    student_name.classList.add("error");
  }

  if (student_age.value === "") {
    student_age.classList.add("error");
  }

  if (student_roll_number.value === "") {
    student_roll_number.classList.add("error");
  }

  if (student_course.value == "Choose Course") {
    student_course.classList.add("error");
  }
  let cookie_token = getCookie("token");
  let createStudent = await fetch(
    `http://localhost:5050/insert-student-details`,
    {
      method: "POST",
      headers: { Token: cookie_token },

      body: JSON.stringify({
        Name: student_name.value,
        Age: parseInt(student_age.value),
        RollNumber: student_roll_number.value,
        ClassesEnrolled: {
          course_name: student_course.value,
        },
      }),
    }
  );
  let response = await createStudent.json();
  if (createStudent.status == 500) {
    error_log.style.visibility = "visible";
    error_log.innerHTML = "Error submitting!<br>" + response;
  } else if (createStudent.status != 500) {
    document.getElementById("responseBody").innerHTML = "Created Student";
    submitStudent.disabled = true;
    submitStudent.classList.add("when_submited");
  }
}

function removeError() {
  let student_name = document.getElementById("name");
  let student_age = document.getElementById("age");
  let student_roll_number = document.getElementById("roll_number");
  let student_course = document.getElementById("cn_drop_down");
  let error_log = document.getElementById("error_log");
  student_name.classList.remove("error");
  student_age.classList.remove("error");
  student_roll_number.classList.remove("error");
  student_course.classList.remove("error");
  error_log.style.visibility = "hidden";
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
