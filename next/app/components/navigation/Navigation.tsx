import { FragmentType, graphql, useFragment } from "@/libs/gql";
import { AutoPlayButton } from "./AutoPlayButton";
import { NextButton } from "./NextButton";
import { PrevButton } from "./PrevButton";
import { StepDisplay } from "./StepDisplay";
import styles from "./style.module.css";

const fragmentDefinition = graphql(/* GraphQL */ `
  fragment Navigation_Fragment on Page {
    step
    nextStep
    prevStep
    durationSeconds
    isTrivialStep
  }
`);

export interface NavigationProps {
  fragment: FragmentType<typeof fragmentDefinition>;
}

export const Navigation = (props: NavigationProps): JSX.Element => {
  const fragment = useFragment(fragmentDefinition, props.fragment);

  return (
    <>
      {fragment.step && <StepDisplay step={fragment.step} />}
      {fragment.prevStep && (
        <PrevButton href={`/?step=${fragment.prevStep}&skipAnimation=true`} />
      )}
      {fragment.nextStep && (
        <AutoPlayButton
          nextStep={fragment.nextStep}
          durationSeconds={fragment.durationSeconds}
          isTrivialStep={fragment.isTrivialStep}
        />
      )}
      {fragment.nextStep && <NextButton href={`/?step=${fragment.nextStep}`} />}
    </>
  );
};
