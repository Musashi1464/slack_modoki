import { Input } from "@chakra-ui/react";
import { ChangeEvent, memo, VFC } from "react";

type Props = {
  placeholder: string;
  value: string;
  onChange: (e: ChangeEvent<HTMLInputElement>) => void;
};

export const FormInput: VFC<Props> = memo((props) => {
  const { placeholder, value, onChange } = props;
  return (
    <Input
      placeholder={placeholder}
      value={value}
      onChange={onChange}
      bg="white"
      borderRadius="xl"
      outline="none"
      m="4px"
      width="40vh"
    />
  );
});
