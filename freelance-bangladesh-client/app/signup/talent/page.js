"use client";

import React from "react";
import { signupTalent } from "@/services/userService";
import Form from "@/components/form";
import { useCanActivePublicComponent } from "@/utils/authorizeHelper";

export default function SignupTalentPage() {
  useCanActivePublicComponent(); // public route guard

  const emailRef = React.useRef();
  const firstnameRef = React.useRef();
  const lastnameRef = React.useRef();
  const phoneRef = React.useRef();
  const passwordRef = React.useRef();
  const confirmPasswordRef = React.useRef();

  const handleCreate = () => {
    if (passwordRef.current.value != confirmPasswordRef.current.value) {
      return alert("Password & confirm password doesn't match!")
    }

    signupTalent({
      email: emailRef.current.value,
      first_name: firstnameRef.current.value,
      last_name: lastnameRef.current.value,
      mobile_number: phoneRef.current.value,
      password: passwordRef.current.value,
      role: "talent",
    }).then(() => {
      alert("signup success!")
      // reset
      emailRef.current.value = null;
      firstnameRef.current.value = null;
      lastnameRef.current.value = null;
      phoneRef.current.value = null;
      passwordRef.current.value = null;
      confirmPasswordRef.current.value = null;
    }).catch((err) => {
      alert(err.message ?? "Some unexpected error has occurred.")
    })
  }

  return (
    <main>
      <Form
        formTitle="Join as a Talent"
        submitBtnName="Create"
        dispatchAction={handleCreate}
        formItems={[
            {
            label: "Email",
            name: "email",
            ref: emailRef,
            type: "text",
            id: "email",
            placeholder: "(a valid email)",
            required: true,
            },
            {
            label: "First Name",
            name: "firstname",
            ref: firstnameRef,
            type: "text",
            id: "firstname",
            placeholder: "Md. Sayeed",
            required: true,
            },
            {
            label: "Last Name",
            name: "lastname",
            ref: lastnameRef,
            type: "text",
            id: "lastname",
            placeholder: "Rahman",
            required: true,
            },
            {
            label: "Mobile Number",
            name: "phone",
            ref: phoneRef,
            type: "text",
            id: "phone",
            placeholder: "+880*******",
            required: true,
            },
            {
            label: "Password",
            name: "password",
            ref: passwordRef,
            type: "password",
            id: "password",
            required: true,
            },
            {
            label: "Confirm Password",
            name: "confirmPassword",
            ref: confirmPasswordRef,
            type: "password",
            id: "confirmPassword",
            required: true,
            },
        ]}
      />
    </main>
  );
}
