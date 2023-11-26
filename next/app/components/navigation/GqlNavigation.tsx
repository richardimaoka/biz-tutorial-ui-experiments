import { FragmentType, graphql, useFragment } from "@/libs/gql";
import styles from "./GqlNavigation.module.css";
import { HandleKeyEvents } from "./HandleKeyEvents";
import { HandleTrivial } from "./HandleTrivial";

const fragmentDefinition = graphql(`
  fragment GqlNavigation on Page {
    prevStep
    nextStep
    isTrivial
  }
`);

interface Props {
  fragment: FragmentType<typeof fragmentDefinition>;
  tutorial: string;
}

export function GqlNavigation(props: Props) {
  const fragment = useFragment(fragmentDefinition, props.fragment);

  return (
    <nav>
      <button className={styles.button}>button</button>
      <HandleKeyEvents
        tutorial={props.tutorial}
        prevStep={fragment.prevStep}
        nextStep={fragment.nextStep}
      />
      <HandleTrivial
        isTrivial={fragment.isTrivial}
        nextStep={fragment.nextStep}
      />
    </nav>
  );
}
