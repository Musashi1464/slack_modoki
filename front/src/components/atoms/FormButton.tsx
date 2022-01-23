import { Button } from "@chakra-ui/react";
import { memo, ReactNode, VFC } from "react";

type Props = {
  children: ReactNode;
  isButtonDisabled: boolean;
  onClick: () => void;
};

export const FormButton: VFC<Props> = memo((props) => {
  const { children, isButtonDisabled, onClick } = props;
  return (
    <Button
      disabled={isButtonDisabled}
      onClick={onClick}
      borderRadius="18px"
      backgroundColor="blue.200"
      paddingX="20px"
      _hover={{ bg: "blue.300" }}
    >
      {children}
    </Button>
  );
});
