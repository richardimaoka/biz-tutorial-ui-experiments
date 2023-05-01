import Link from "next/link";

interface NextStepButtonProps {
  nextStep: string;
  openFilePath?: string;
}

export const NextStepButton = ({
  nextStep,
  openFilePath,
}: NextStepButtonProps): JSX.Element => {
  const queryString = openFilePath
    ? `step=${nextStep}&openFilePath=${encodeURIComponent(openFilePath)}`
    : `step=${nextStep}`;

  return (
    <button type="button">
      <Link href={`./?${queryString}`}>next step</Link>
    </button>
  );
};
