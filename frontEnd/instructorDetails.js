async function populateInstructors() {
  let cookie_token = getCookie("token");
  let all_instructors = await fetch(
    `http://localhost:5050/retrieve-instructors`,
    {
      credentials: "same-origin",
      headers: { token: cookie_token },
    }
  );
  let all_instructors_response = await all_instructors.json();
  if (all_instructors_response == "authentication time-out") {
    alert("Timed-out re login");
    setTimeout(window.location.replace("index.html"), 2000);
  }
  for (let i = 0; i < all_instructors_response.length; i++) {
    let each_value = all_instructors_response[i];
    let table1 = document.getElementById("instructor_table");
    let tr = document.createElement("tr");
    tr.innerHTML = `<td id=${each_value.instructor_code} >${
      each_value.instructor_code
    }</td>
       <td id=${i}>${each_value.instructor_name}</td>
       <td id=${each_value.department.replace(" ", "") + i} >${
      each_value.department
    }</td>
       <td id=${each_value.course_name[0] + i}>${each_value.course_name}</td>
       <td><button onclick=openForm("${each_value.instructor_code}",${i},${
      each_value.department.replace(" ", "") + i
    },${each_value.course_name[0] + i}) class="update_button">U</button></td>
       <td><button onclick=deleteInstructor(${i},${
      each_value.course_name[0] + i
    }) class="delete_button">X</button></td>`;
    table1.appendChild(tr);
  }
}
populateInstructors();
function setdashboard() {
  window.location.replace("allinstructor.html");
}
function setbackpage() {
  window.location.replace("createInstructor.html");
}
async function deleteInstructor(index, course_name_index) {
  let index_name = document.getElementById(index);
  let cookie_token = getCookie("token");
  let deleteCourse = await fetch(`http://localhost:5050/delete-instructor`, {
    method: "DELETE",
    headers: { Token: cookie_token },
    body: JSON.stringify({
      instructor_name: index_name.innerText,
      course_name: course_name_index.innerText,
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

async function UpdateInstructor(code, name_index, dpt, course_name) {
  let index_name = document.getElementById(name_index);
  let cookie_token = getCookie("token");
  let old_instructor_code = document.getElementById("old_instructor_code");
  let old_instructor_name = document.getElementById("old_instructor_name");
  let old_department_name = document.getElementById("old_department_name");
  let old_course_name = document.getElementById("old_course_name");

  url = `http://localhost:5050/update-instructor`;
  if (old_instructor_code.innerHTML != "") {
    url = url + `?instructor_code=${old_instructor_code.innerHTML}`;
  }
  if (old_instructor_name.innerHTML != "") {
    url = url + `&instructor_name=${old_instructor_name.innerHTML}`;
  }
  if (old_course_name.innerHTML != "") {
    url = url + `&course_name=${old_course_name.innerHTML}`;
  }
  let updateCourse = await fetch(url, {
    method: "PATCH",
    headers: { Token: cookie_token },
    body: JSON.stringify({
      instructor_code: code,
      instructor_name: name_index,
      department: dpt,
      course_name: course_name,
    }),
  });
  let response = await updateCourse.json();
  if (!updateCourse.ok) {
    console.log("failed", response);
  } else {
    console.log("success", response);
    window.location.reload();
  }
}
function openForm(code, name_index, dpt, course_name) {
  let popup = document.getElementById("popup");
  popup.classList.add("open-popup");

  let index_name = document.getElementById(name_index);
  let old_instructor_code = document.getElementById("old_instructor_code");
  let old_instructor_name = document.getElementById("old_instructor_name");
  let old_department_name = document.getElementById("old_department_name");
  let old_course_name = document.getElementById("old_course_name");

  old_instructor_code.innerHTML = code;
  old_instructor_name.innerHTML = index_name.innerHTML;
  old_department_name.innerHTML = dpt.innerHTML;
  old_course_name.innerHTML = course_name.innerHTML;
}
function closeForm() {
  let popup = document.getElementById("popup");
  popup.classList.remove("open-popup");
}
function callUpdateFunction() {
  let req_instructor_code = document.getElementById("req_instructor_code");
  let req_instructor_name = document.getElementById("req_instructor_name");
  let req_department_name = document.getElementById("req_department_name");
  let req_course_name = document.getElementById("req_course_name");

  UpdateInstructor(
    req_instructor_code.value,
    req_instructor_name.value,
    req_department_name.value,
    req_course_name.value
  );
}
