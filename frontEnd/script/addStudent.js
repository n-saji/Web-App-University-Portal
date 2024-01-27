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
    `http://3.111.149.112:5050/insert-student-details`,
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
