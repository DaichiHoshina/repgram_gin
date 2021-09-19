import { Button, Card } from "@material-ui/core";
import { FC } from "react";
import { useDispatch } from "react-redux";
import { signUpUser } from "../../../services/User";
import TextFieldParts from "../../atoms/share/TextFieldParts";
import { useRouter } from "next/router";
import KeyValueColonPair from "../common/KeyValueColonPair";
import { useFormik } from "formik";
import { useSnackbar } from "notistack";
import { AccountCreateSchema } from "../../../const/validation";
interface Props {
  isSubmit: boolean;
}

const SignUpForm: FC<Props> = () => {
  const router = useRouter();
  const dispatch = useDispatch();
  const { enqueueSnackbar } = useSnackbar() || {};

  const formik = useFormik({
    initialValues: { name: "", email: "", password: "" },
    validationSchema: AccountCreateSchema,
    onSubmit: async (values) => {
      const result: any = await dispatch(signUpUser({ loginForm: values }));
      if (result.payload.status === 200) {
        router.push("/presentations");
        enqueueSnackbar("ユーザー登録しました", {
          variant: "success",
        });
      } else {
        enqueueSnackbar("ユーザー登録に失敗しました", {
          variant: "error",
        });
      }
    },
  });

  return (
    <Card
      className="flex-item flex flex-col px-8 py-12 w-1/2 max-w-5xl"
      style={{ color: "#ffffff", backgroundColor: "#242323" }}
    >
      <h1 className="pb-8 text-center text-xl">Sign Up</h1>
      <form action="" onSubmit={formik.handleSubmit}>
        <ul className="pb-10 space-y-3">
          <KeyValueColonPair
            keyName="氏名"
            value={<TextFieldParts name="name" formik={formik} />}
          />
          <KeyValueColonPair
            keyName="メールアドレス"
            value={<TextFieldParts name="email" formik={formik} />}
          />
          <KeyValueColonPair
            keyName="パスワード"
            value={
              <TextFieldParts
                name="password"
                isPasswordForm={true}
                formik={formik}
              />
            }
          />
        </ul>
        <Button
          type="submit"
          data-testid="form"
          variant="contained"
          color="primary"
          disabled={
            !formik.values.name ||
            !formik.values.email ||
            !formik.values.password
          }
          fullWidth
          className="border-none ring-transparent"
        >
          Sign Up
        </Button>
      </form>
    </Card>
  );
};
export default SignUpForm;
