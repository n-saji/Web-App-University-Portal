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
  }
  for (let i = 0; i < all_students_response.length; i++) {
    let each_value = all_students_response[i];
    let table1 = document.getElementById("student_table");
    let tr = document.createElement("tr");
    tr.innerHTML = `<td>${each_value.Name}</td>
       <td id=${i}>${each_value.RollNumber}</td>
       <td>${each_value.Age}</td>
       <td id=${each_value.ClassesEnrolled.course_name[0] + i}>${
      each_value.ClassesEnrolled.course_name
    }</td>
       <td>${each_value.StudentMarks.Marks}</td>
       <td>${each_value.StudentMarks.Grade}</td>
       <td><button class="update_button">U</button></td>
       <td><button onclick=deleteStudent(${i},${
      each_value.ClassesEnrolled.course_name[0] + i
    }) class="delete_button">X</button></td>`;
    table1.appendChild(tr);
  }
}
populateInstructors();
