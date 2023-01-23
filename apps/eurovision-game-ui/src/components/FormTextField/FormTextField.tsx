import { TextField, TextFieldProps } from "@mui/material";
import { Field, FormikErrors, FormikTouched } from "formik";

type TFormFieldProps<T> = TextFieldProps & {
	name: keyof T;
	errors?: FormikErrors<T>;
	touched?: FormikTouched<T>;
};

const FormTextField = <T extends object>({
	name,
	errors,
	touched,
	...rest
}: TFormFieldProps<T>) => {
	const isError = Boolean(errors?.[name] && touched?.[name]);
	const helperText = errors?.[name];
	return (
		<Field
			name={name}
			error={isError}
			helperText={isError && helperText}
			{...rest}
			as={TextField}
		/>
	);
};

export default FormTextField;
