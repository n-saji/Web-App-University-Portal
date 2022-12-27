function userlogin() {
  const emailId = document.getElementById("username").value;
  if (!emailId) {
    alert("Please enter username.");
    return;
  }
  const password = document.getElementById("password").value;
  if (!password) {
    alert("Please enter password.");
    return;
  }
  const returnvalue = document.getElementById("loginButton");

  returnvalue.addEventListener("click", toCheckValidity(emailId, password));
}

function forgotPassword() {
  const returnvalue = document.getElementById("tempfix");
  returnvalue.style.display = "block";
  returnvalue.innerHTML = "&#9888Feature under development&#9888";
  setTimeout(disablefunction, 3000);
}
function disablefunction() {
  const returnvalue = document.getElementById("tempfix");
  returnvalue.style.display = "none";
}

async function toCheckValidity(emailId, password) {
  let uuid = await CheckValidity(emailId, password).catch((error) =>
    console.log(error)
  );
  if (uuid != "") {
    console.log("correct", uuid);
    window.location.replace("allInstructor.html");
  }
}

async function CheckValidity(username, password) {
  let response = await fetch(
    `http://localhost:5050/instructor-login/${username}/${password}`
  );
  let uuid_instructor = await response.json();
  if (response.ok === true) {
    // window
    //   .open(
    //     `http://localhost:5050/retrieve-all-courses/${responsemsg}`,
    //     "_blank"
    //   )
    //   .focus();
    // window.location.replace(
    //   `http://localhost:5050/retrieve-all-courses/${responsemsg}`
    // );
    //window.location.replace("allCourse.html");
    return uuid_instructor;
  } else {
    alert("HTTP-Error: " + uuid_instructor);
    return "";
  }
}
