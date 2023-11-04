import { GqlFileNodeIcon } from "./GqlFileNodeIcon";
import styles from "./style.module.css";

import { FragmentType, graphql, useFragment } from "@/libs/gql";
import { LinkSearchParams } from "@/app/components/link/LinkSearchParams";

const fragmentDefinition = graphql(`
  fragment GqlFileNodeComponent on FileNode {
    ...GqlFileNodeIcon
    nodeType
    name
    filePath
    offset
    isUpdated
  }
`);

interface Props {
  fragment: FragmentType<typeof fragmentDefinition>;
  currentDirectory?: string;
  step: string;
}

export function GqlFileNodeComponent(props: Props) {
  const fragment = useFragment(fragmentDefinition, props.fragment);

  const isCurrentDirectory =
    fragment.filePath &&
    props.currentDirectory &&
    fragment.filePath === props.currentDirectory;

  const filenodeStyle = fragment.isUpdated
    ? `${styles.component} ${styles.updated}`
    : styles.component;

  const offset = fragment.offset ? fragment.offset : 0;

  return (
    <LinkSearchParams searchParams={{ file: fragment.filePath }}>
      <div className={filenodeStyle}>
        <div style={{ marginLeft: `${offset * 16}px` }}>
          <GqlFileNodeIcon fragment={fragment} />
        </div>
        <div className={styles.nodename}>
          {fragment.name}
          {isCurrentDirectory && (
            <span className={styles.currentdirectory}>
              {" //current directory"}
            </span>
          )}
        </div>
      </div>
    </LinkSearchParams>
  );
}
