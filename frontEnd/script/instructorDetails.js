async function populateInstructors() {
  let cookie_token = getCookie("token");
  let all_instructors = await fetch(
    `http://3.111.149.112:5050/retrieve-instructors`,
    {
      credentials: "same-origin",
      headers: { token: cookie_token },
    }
  );
  let all_instructors_response = await all_instructors.json();
  for (let i = 0; i < all_instructors_response.length; i++) {
    let each_value = all_instructors_response[i];
    let table_body = document.getElementById("table_body");
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

    table_body.appendChild(tr);
  }
}
populateInstructors();
function setdashboard() {
  window.location.replace("dashboard-v2.html");
}
function setbackpage() {
  window.location.replace("createInstructor.html");
}

async function deleteInstructor(index, course_name_index, department) {
  let index_name = document.getElementById(index);
  let cookie_token = getCookie("token");
  let deleteCourse = await fetch(
    `http://3.111.149.112:5050/delete-instructor`,
    {
      method: "DELETE",
      headers: { Token: cookie_token },
      body: JSON.stringify({
        instructor_name: index_name.innerText,
        course_name: course_name_index.innerText,
      }),
    }
  );

  let response = await deleteCourse.json();
  if (!deleteCourse.ok) {
    console.log("failed", response);
  } else {
    console.log("success", response);
    window.location.reload();
  }
}

async function UpdateInstructor(code, name_index, dpt, course_name) {
  let cookie_token = getCookie("token");
  let old_instructor_code = document.getElementById("old_instructor_code");
  let old_instructor_name = document.getElementById("old_instructor_name");
  let old_course_name = document.getElementById("old_course_name");
  let popup = document.getElementById("popup");

  let url = `http://3.111.149.112:5050/update-instructor`;
  if (old_instructor_code.innerHTML != "") {
    url = url + `?instructor_code=${old_instructor_code.innerHTML}`;
  }
  if (old_instructor_name.innerHTML != "") {
    url = url + `&instructor_name=${old_instructor_name.innerHTML}`;
  }
  if (old_course_name.innerHTML != "") {
    url = url + `&course_name=${old_course_name.innerHTML}`;
  }

  if (course_name == "Choose Course") {
    course_name = "";
  }

  let updateInstructor = await fetch(url, {
    method: "PATCH",
    headers: { Token: cookie_token },
    body: JSON.stringify({
      instructor_code: code,
      instructor_name: name_index,
      department: dpt,
      course_name: course_name,
    }),
  });
  let response = await updateInstructor.json();
  if (!updateInstructor.ok) {
    let err = document.getElementById("error_msg");
    err.classList.add("err_msg");
    err.innerHTML = response;
    console.log("failed", response);
    popup.scrollTo(0, 600);
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

window.onkeydown = function (event) {
  if (event.keyCode == 27) {
    closeForm();
  }
};

function closeForm() {
  let popup = document.getElementById("popup");
  let err = document.getElementById("error_msg");
  err.innerHTML = "";
  popup.classList.remove("open-popup");
}
function callUpdateFunction() {
  let req_instructor_code = document.getElementById("req_instructor_code");
  let req_instructor_name = document.getElementById("req_instructor_name");
  let req_department_name = document.getElementById("req_department_name");
  let req_course_name = document.getElementById("cn_drop_down");

  UpdateInstructor(
    req_instructor_code.value,
    req_instructor_name.value,
    req_department_name.value,
    req_course_name.value
  );
}

async function populateInstructorsWithCondition(order) {
  let cookie_token = getCookie("token");
  let all_instructors = await fetch(
    `http://3.111.149.112:5050/retrieve-instructors/${order}`,
    {
      headers: { token: cookie_token },
    }
  );
  let all_instructors_response = await all_instructors.json();

  let table_body = document.getElementById("table_body");
  table_body.innerText = "";
  for (let i = 0; i < all_instructors_response.length; i++) {
    let each_value = all_instructors_response[i];

    let tr = document.createElement("tr");
    tr.innerHTML = `<td id=${each_value.instructor_code} >${
      each_value.instructor_code
    }</td>
       <td id=${i}>${each_value.instructor_name}</td>
       <td id=${each_value.department.replace(" ", "") + i} >${
      each_value.department
    }</td>
       <td id=${each_value.course_name[0] + i}>${
      each_value.ClassesEnrolled.course_name
    }</td>
       <td><button onclick=openForm("${each_value.instructor_code}",${i},${
      each_value.department.replace(" ", "") + i
    },${each_value.course_name[0] + i}) class="update_button">U</button></td>
       <td><button onclick=deleteInstructor(${i},${
      each_value.course_name[0] + i
    }) class="delete_button">X</button></td>`;

    table_body.appendChild(tr);
  }
}
