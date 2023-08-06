import { AnglesLeftIcon } from "@/app/components/icons/AnglesLeftIcon";
import { AnglesRightIcon } from "@/app/components/icons/AnglesRightIcon";
import styles from "./style.module.css";

import { FragmentType, graphql, useFragment } from "@/libs/gql";

const fragmentDefinition = graphql(`
  fragment FileTreeHeader_Fragment on SourceCode {
    projectDir
  }
`);

interface FileTreeHeaderProps {
  fragment: FragmentType<typeof fragmentDefinition>;
  isFolded: boolean;
  onButtonClick: () => void;
}

export const FileTreeHeader = (props: FileTreeHeaderProps): JSX.Element => {
  const fragment = useFragment(fragmentDefinition, props.fragment);

  const headerStyle = props.isFolded
    ? `${styles.header} ${styles.folded}`
    : `${styles.header} ${styles.expanded}`;

  const ProjectDir = () =>
    fragment.projectDir ? (
      <div className={styles.projectdir}>{fragment.projectDir}</div>
    ) : (
      <div />
    );

  return (
    <div className={headerStyle}>
      {!props.isFolded && <ProjectDir />}
      <button onClick={props.onButtonClick}>
        {props.isFolded ? <AnglesRightIcon /> : <AnglesLeftIcon />}
      </button>
    </div>
  );
};
