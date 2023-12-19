import { FragmentType, graphql, useFragment } from "@/libs/gql";
import { ButtonToInitialStep } from "./buttons/ButtonToInitialStep";
import { GqlColumnTabs } from "./tabs/GqlColumnTabs";
import styles from "./GqlHandsonHeader.module.css";

const fragmentDefinition = graphql(`
  fragment GqlHandsonHeader on Page {
    ...GqlColumnTabs
  }
`);

interface Props {
  fragment: FragmentType<typeof fragmentDefinition>;
}

export function GqlHandsonHeader(props: Props) {
  const fragment = useFragment(fragmentDefinition, props.fragment);

  return (
    <div className={styles.component}>
      <GqlColumnTabs fragment={fragment} />
      <ButtonToInitialStep />
    </div>
  );
}
