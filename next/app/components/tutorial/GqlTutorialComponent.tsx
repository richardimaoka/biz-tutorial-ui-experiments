import { FragmentType, graphql, useFragment } from "@/libs/gql";
import styles from "./GqlTutorialComponent.module.css";
import { GqlColumnWrappers } from "./column/GqlColumnWrappers";
import { GqlTutorialHeader } from "./header/GqlTutorialHeader";

const fragmentDefinition = graphql(`
  fragment GqlTutorialComponent on Page2 {
    ...GqlTutorialHeader
    ...GqlColumnWrappers
  }
`);

interface Props {
  fragment: FragmentType<typeof fragmentDefinition>;
  selectTab: string;
}

export function GqlTutorialComponent(props: Props) {
  const fragment = useFragment(fragmentDefinition, props.fragment);

  return (
    <div className={styles.component}>
      {/* header part */}
      <div className={styles.header}>
        <GqlTutorialHeader fragment={fragment} selectTab={props.selectTab} />
      </div>
      {/* contents part */}
      <div className={styles.contents}>
        <GqlColumnWrappers fragment={fragment} />
      </div>
    </div>
  );
}
