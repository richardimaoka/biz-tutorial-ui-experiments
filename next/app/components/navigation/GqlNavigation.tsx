import { FragmentType, graphql, useFragment } from "@/libs/gql";
import styles from "./GqlNavigation.module.css";
import { HandleKeyEvents } from "./HandleKeyEvents";
import { HandleTrivial } from "./HandleTrivial";
import { NextButton } from "./NextButton";
import { PrevButton } from "./PrevButton";

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

  // If isTrivial, user cannnot control prev/next navigatio
  const allowUserNavigation = !fragment.isTrivial;

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
      {allowUserNavigation && fragment.prevStep && (
        <PrevButton prevStep={fragment.prevStep} />
      )}
      {allowUserNavigation && fragment.nextStep && (
        <NextButton nextStep={fragment.nextStep} />
      )}
      {allowUserNavigation && (
        <HandleKeyEvents
          tutorial={props.tutorial}
          prevStep={fragment.prevStep}
          nextStep={fragment.nextStep}
        />
      )}
      <HandleTrivial
        isTrivial={fragment.isTrivial}
        nextStep={fragment.nextStep}
      />
    </nav>
  );
}
