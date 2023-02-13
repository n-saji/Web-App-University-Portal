function setbackpage() {
  window.location.replace("addStudent.html");
}

function setdashboard() {
  window.location.replace("allinstructor.html");
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
async function deleteStudent(index, course_name_index) {
  let index_name = document.getElementById(index);
  let cookie_token = getCookie("token");
  let deleteCourse = await fetch(`http://localhost:5050/delete-student`, {
    method: "DELETE",
    headers: { Token: cookie_token },
    body: JSON.stringify({
      RollNumber: index_name.innerText,
      ClassesEnrolled: {
        course_name: course_name_index.innerText,
      },
    }),
  });
  let response = await deleteCourse.json();
  if (!deleteCourse.ok) {
    console.log("failed", response);
  } else {
    console.log("success", response);
    window.location.reload();
  }
}
async function populateInstructors() {
  let cookie_token = getCookie("token");
  let all_students = await fetch(
    `http://localhost:5050/retrieve-college-administration`,
    {
      credentials: "same-origin",
      headers: { token: cookie_token },
    }
  );
  let all_students_response = await all_students.json();
  if (all_students_response == "token expired! Generate new token") {
    alert("Timed-out re login");
    setTimeout(window.location.replace("index.html"), 2000);
    return;
  }
  for (let i = 0; i < all_students_response.length; i++) {
    let each_value = all_students_response[i];
    let table1 = document.getElementById("student_table");
    let tr = document.createElement("tr");
    tr.innerHTML = `<td id=${i}>${each_value.RollNumber}</td>
        <td id=${each_value.Name.replace(" ", "") + i}>${each_value.Name}</td>
        <td id=${each_value.Age}>${each_value.Age}</td>
        <td id=${each_value.ClassesEnrolled.course_name[0] + i}>${
      each_value.ClassesEnrolled.course_name
    }</td>
        <td id=${each_value.StudentMarks.Marks}>${
      each_value.StudentMarks.Marks
    }</td>
        <td>${each_value.StudentMarks.Grade}</td>
        <td><button onclick=openPopUpByUpdate(${i},${
      each_value.Name.replace(" ", "") + i
    },${each_value.Age},${each_value.ClassesEnrolled.course_name[0] + i},${
      each_value.StudentMarks.Marks
    }) class="update_button">U</button></td>
        <td><button onclick=deleteStudent(${i},${
      each_value.ClassesEnrolled.course_name[0] + i
    }) class="delete_button">X</button></td>`;
    table1.appendChild(tr);
  }
}
populateInstructors();

function openPopUpByUpdate(roll_number, name, age, course_name, marks) {
  let old_roll_number_inner_html = document.getElementById(roll_number);

  let popup = document.getElementById("popup");
  popup.classList.add("open-popup");

  let old_roll_number = document.getElementById("old_roll_number");
  let old_name = document.getElementById("old_name");
  let old_age = document.getElementById("old_age");
  let old_course_name = document.getElementById("old_course_name");
  let old_marks = document.getElementById("old_marks");

  old_roll_number.innerHTML = old_roll_number_inner_html.innerHTML;
  old_name.innerHTML = name.innerHTML;
  old_age.innerHTML = age;
  old_course_name.innerHTML = course_name.innerHTML;
  old_marks.innerHTML = marks;
}
async function updateStudent(
  req_roll_number,
  req_name,
  req_age,
  req_course_name,
  req_marks
) {
  let old_roll_number = document.getElementById("old_roll_number");
  let old_name = document.getElementById("old_name");
  let old_age = document.getElementById("old_age");
  let old_course_name = document.getElementById("old_course_name");
  let old_marks = document.getElementById("old_marks");
  if (req_roll_number == "") {
    req_roll_number = old_roll_number.innerHTML;
  }
  if (req_name == "") {
    req_name = old_name.innerHTML;
  }
  if (req_age == 0) {
    req_age = old_age.innerHTML;
  }
  if (req_course_name == "Choose Course") {
    req_course_name = old_course_name.innerHTML;
  }
  if (req_marks == 0) {
    req_marks = old_marks.innerHTML;
  }
  let cookie_token = getCookie("token");
  let updateStudent = await fetch(
    `http://localhost:5050/update-student-details/${old_course_name.innerHTML}`,
    {
      method: "PATCH",
      headers: { Token: cookie_token },
      body: JSON.stringify({
        Name: req_name,
        RollNumber: req_roll_number,
        Age: parseInt(req_age),
        ClassesEnrolled: {
          course_name: req_course_name,
        },
        StudentMarks: {
          Marks: parseInt(req_marks),
        },
      }),
    }
  );
  let response = await updateStudent.json();
  if (updateStudent.status != 200) {
    console.log("failed", response);
  } else {
    console.log("success", response);
    window.location.reload();
  }
  let popup = document.getElementById("popup");
  popup.classList.remove("open-popup");
}

function closePopUpByCancel() {
  let popup = document.getElementById("popup");
  popup.classList.remove("open-popup");
}
function closePopUpBySubmit() {
  let req_roll_number = document.getElementById("pop_up_roll_number");
  let req_name = document.getElementById("pop_up_name");
  let req_age = document.getElementById("pop_up_age");
  let req_course_name = document.getElementById("cn_drop_down");
  let req_marks = document.getElementById("pop_up_marks");
  console.log(
    req_roll_number.value,
    req_name.value,
    req_age.value,
    req_course_name.value,
    req_marks.value
  );
  updateStudent(
    req_roll_number.value,
    req_name.value,
    req_age.value,
    req_course_name.value,
    req_marks.value
  );
}

window.onkeydown = function (event) {
  if (event.keyCode == 27) {
    closePopUpByCancel();
  }
};

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
