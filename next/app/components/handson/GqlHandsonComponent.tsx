import { FragmentType, graphql, useFragment } from "@/libs/gql";
import styles from "./GqlHandsonComponent.module.css";
import { GqlColumnWrappers } from "./column/GqlColumnWrappers";
import { GqlHandsonHeader } from "./header/GqlHandsonHeader";

const fragmentDefinition = graphql(`
  fragment GqlHandsonComponent on Page {
    ...GqlHandsonHeader
    ...GqlColumnWrappers
  }
`);

interface Props {
  fragment: FragmentType<typeof fragmentDefinition>;
}

export function GqlHandsonComponent(props: Props) {
  const fragment = useFragment(fragmentDefinition, props.fragment);

  return (
    <div className={styles.component}>
      {/* header part */}
      <div className={styles.header}>
        <GqlHandsonHeader fragment={fragment} />
      </div>
      {/* contents part */}
      <div className={styles.contents}>
        <GqlColumnWrappers fragment={fragment} />
      </div>
    </div>
  );
}
