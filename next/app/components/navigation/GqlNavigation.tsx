import { FragmentType, graphql, useFragment } from "@/libs/gql";
import styles from "./GqlNavigation.module.css";
import { HandleKeyEvents } from "./HandleKeyEvents";
import { HandleTrivial } from "./HandleTrivial";
import { NextButton } from "./NextButton";

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
    <nav
    // style={{
    //   position: "absolute",
    //   backgroundColor: "white",
    //   width: "100%",
    //   height: "100%",
    //   zIndex: 200,
    //   top: "0%",
    // }}
    >
      {fragment.nextStep && <NextButton nextStep={fragment.nextStep} />}
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
