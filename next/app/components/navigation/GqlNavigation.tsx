import { FragmentType, graphql, useFragment } from "@/libs/gql";
import styles from "./GqlNavigation.module.css";
import { HandleKeyEvents } from "./HandleKeyEvents";
import { HandleTrivial } from "./HandleTrivial";
import { NextButton } from "./NextButton";
import { PrevButton } from "./PrevButton";
import { InitPageButton } from "./InitPageButton";

const fragmentDefinition = graphql(`
  fragment GqlNavigation on Page {
    prevStep
    nextStep
    isTrivial
  }
`);

interface Props {
  fragment: FragmentType<typeof fragmentDefinition>;
  toInitial?: boolean;
}

export function GqlNavigation(props: Props) {
  const fragment = useFragment(fragmentDefinition, props.fragment);

  // If isTrivial, user cannnot control prev/next navigatio
  const disableUserNavigation = fragment.isTrivial === true;

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
      {props.toInitial && <InitPageButton />}
      {fragment.prevStep && (
        <PrevButton
          prevStep={fragment.prevStep}
          disabled={disableUserNavigation}
        />
      )}
      {fragment.nextStep && (
        <NextButton
          nextStep={fragment.nextStep}
          disabled={disableUserNavigation}
        />
      )}
      {!disableUserNavigation && (
        <HandleKeyEvents
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
