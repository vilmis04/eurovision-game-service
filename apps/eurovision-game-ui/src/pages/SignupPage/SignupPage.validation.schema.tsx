import * as Yup from "yup";

export const signupValidationSchema = Yup.object().shape({
	username: Yup.string().required("Required"),
	password: Yup.string().required("Required"),
	repeatPassword: Yup.string()
		.required("Required")
		.oneOf([Yup.ref("password")], "Passwords must match"),
});
