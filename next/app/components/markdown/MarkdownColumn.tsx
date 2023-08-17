import { FragmentType, graphql, useFragment } from "@/libs/gql";
import { MarkdownView } from "../markdown/MarkdownView";
import styles from "./style.module.css";

const fragmentDefinition = graphql(`
  fragment MarkdownColumn_Fragment on MarkdownColumn {
    description {
      ...MarkdownFragment
    }
    contentsPosition
  }
`);

export interface MarkdownColumnProps {
  fragment: FragmentType<typeof fragmentDefinition>;
}

export const MarkdownColumn = (props: MarkdownColumnProps): JSX.Element => {
  const fragment = useFragment(fragmentDefinition, props.fragment);

  const positionStyle = () => {
    switch (fragment.contentsPosition) {
      case "TOP":
        return styles.verticalTop;
      case "CENTER":
        return styles.verticalCenter;
      case "BOTTOM":
        return styles.verticalBottom;
      default:
        return styles.verticalCenter;
    }
  };

  if (!fragment.description) return <></>;

  return (
    <div className={`${styles.column} ${positionStyle()}`}>
      <MarkdownView fragment={fragment.description} />
    </div>
  );
};
