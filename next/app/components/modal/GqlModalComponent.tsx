import { FragmentType, graphql, useFragment } from "@/libs/gql";
import { MarkdownConfigurable } from "../markdown/server-component/MarkdownConfigurable";
import styles from "./GqlModalComponent.module.css";
import modalMdStyles from "./ModalMarkdown.module.css";

const fragmentDefinition = graphql(`
  fragment GqlModalComponent on Modal {
    markdownBody
    position
  }
`);

interface Props {
  fragment: FragmentType<typeof fragmentDefinition>;
}

export function GqlModalComponent(props: Props) {
  const fragment = useFragment(fragmentDefinition, props.fragment);

  const position = fragment.position ? fragment.position : "TOP";
  let positionStyle: string;
  switch (position) {
    case "TOP":
      positionStyle = styles.top;
      break;
    case "CENTER":
      positionStyle = styles.center;
      break;
    case "BOTTOM":
      positionStyle = styles.bottom;
      break;
  }
  const componentStyle = styles.component + " " + positionStyle;

  return (
    <div className={componentStyle}>
      <div className={styles.inner}>
        <MarkdownConfigurable
          className={modalMdStyles.modalMarkdown}
          markdownBody={fragment.markdownBody}
        />
      </div>
    </div>
  );
}
