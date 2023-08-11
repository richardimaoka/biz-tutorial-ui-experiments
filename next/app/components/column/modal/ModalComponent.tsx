import { FragmentType, graphql, useFragment } from "@/libs/gql";
import { ModalPosition } from "@/libs/gql/graphql";
import styles from "./style.module.css";

const fragmentDefinition = graphql(`
  fragment ModalComponentFragment on Modal {
    text
    position
  }
`);

const ModalBox = ({ message }: { message: string }) => (
  <div className={styles.box}>{message}</div>
);

const positionStyle = (p: ModalPosition): string => {
  switch (p) {
    case "TOP":
      return styles.top;
    case "CENTER":
      return styles.center;
    case "BOTTOM":
      return styles.bottom;
  }
};

interface ModalComponentProps {
  fragment: FragmentType<typeof fragmentDefinition>;
}

export const ModalComponent = (props: ModalComponentProps): JSX.Element => {
  const fragment = useFragment(fragmentDefinition, props.fragment);
  const stylePos = fragment.position
    ? positionStyle(fragment.position)
    : positionStyle("CENTER"); //default position = CENTER

  if (!fragment.text) {
    return <></>;
  }

  return (
    <div className={`${stylePos} ${styles.boxposition}`}>
      <ModalBox message={fragment.text} />
    </div>
  );
};
