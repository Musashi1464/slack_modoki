import { Box, Flex, Heading, Divider, Text } from "@chakra-ui/react";
import { ChangeEvent, memo, useState, VFC } from "react";
import { useNavigate } from "react-router-dom";

import { FormButton } from "../atoms/FormButton";
import { FormInput } from "../atoms/FormInput";
import { useAxiosCreate } from "../hooks/useAxiosCreate";

export const CreateWorkspace: VFC = memo(() => {
  const [workspaceName, setWorkspaceName] = useState("");
  const [userName, setUserName] = useState("");
  const [email, setEmail] = useState("");
  const [password, setPassword] = useState("");
  const [isButtonDisabled, setIsButtonDisabled] = useState(true);
  const [isMailDisplay, setIsMainDisplay] = useState("none");
  const navigate = useNavigate();
  const instance = useAxiosCreate();

  const checkForm = (
    valueWorkspaceName: string,
    valueUserName: string,
    valueEmail: string,
    valuePassword: string
  ) => {
    if (
      valueWorkspaceName !== "" &&
      valueUserName !== "" &&
      valueEmail !== "" &&
      valuePassword !== ""
    ) {
      setIsButtonDisabled(false);
    } else {
      setIsButtonDisabled(true);
    }
  };

  const onChangeWorkspaceName = (e: ChangeEvent<HTMLInputElement>) => {
    const valueWorkspaceName = e.target.value;
    setWorkspaceName(valueWorkspaceName);
    checkForm(valueWorkspaceName, userName, email, password);
  };
  const onChangeUserName = (e: ChangeEvent<HTMLInputElement>) => {
    const valueUserName = e.target.value;
    setUserName(valueUserName);
    checkForm(workspaceName, valueUserName, email, password);
  };
  const onChangeEmail = (e: ChangeEvent<HTMLInputElement>) => {
    const valueEmail = e.target.value;
    setEmail(valueEmail);
    checkForm(workspaceName, userName, valueEmail, password);
    const regex =
      /^[a-zA-Z0-9.!#$%&'*+\/=?^_`{|}~-]+@[a-zA-Z0-9-]+(?:\.[a-zA-Z0-9-]+)*$/;
    if (regex.test(valueEmail)) {
      setIsMainDisplay("none");
      setIsButtonDisabled(false);
    } else {
      setIsMainDisplay("");
      setIsButtonDisabled(true);
    }
  };
  const onChangePassword = (e: ChangeEvent<HTMLInputElement>) => {
    const valuePassword = e.target.value;
    setPassword(valuePassword);
    checkForm(workspaceName, userName, email, valuePassword);
  };

  const onClickCreate = () => {
    if (
      workspaceName === "" ||
      userName === "" ||
      email === "" ||
      password === ""
    ) {
      alert("正しくフォームに入力してください。");
    } else {
      instance
        .post("/create_workspace", {
          WorkspaceName: workspaceName,
          UserName: userName,
          Email: email,
          Password: password,
        })
        .then((res) => {
          console.log(res.data);
          navigate("/login");
        })
        .catch((err) => {
          alert("signup failed" + err);
        })
        .finally(() => {
          setWorkspaceName("");
          setUserName("");
          setEmail("");
          setPassword("");
          setIsButtonDisabled(true);
        });
    }
  };

  return (
    <Flex align="center" justify="center" height="80vh">
      <Box
        width="450px"
        bg="white"
        padding="30px"
        borderRadius="xl"
        shadow="md"
      >
        <Heading align="center" pb="10px">
          ワークスペースの作成
        </Heading>
        <Divider />
        <Box align="center" pt="10px">
          <FormInput
            placeholder="ワークスペース名"
            value={workspaceName}
            onChange={onChangeWorkspaceName}
          />
          <FormInput
            placeholder="ユーザー名"
            value={userName}
            onChange={onChangeUserName}
          />
          <FormInput
            placeholder="Eメール"
            value={email}
            onChange={onChangeEmail}
          />
          <Text
            display={isMailDisplay}
            bg="gray.200"
            color="red.300"
            borderRadius="18px"
            py="5px"
            width="40vh"
          >
            ＊正しいパスワードを入力してください。
          </Text>
          <FormInput
            placeholder="パスワード"
            value={password}
            onChange={onChangePassword}
          />
        </Box>
        <Box align="center" pt="10px">
          <FormButton
            isButtonDisabled={isButtonDisabled}
            onClick={onClickCreate}
          >
            作成
          </FormButton>
        </Box>
      </Box>
    </Flex>
  );
});
