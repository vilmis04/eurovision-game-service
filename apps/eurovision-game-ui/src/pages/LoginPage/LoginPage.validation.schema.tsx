import * as Yup from "yup";

export const loginValidationSchema = Yup.object().shape({
	username: Yup.string().required("Required"),
	password: Yup.string().required("Required"),
});
