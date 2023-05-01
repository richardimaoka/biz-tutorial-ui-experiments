import Link from "next/link";

interface PrevStepButtonProps {
  prevStep: string;
  openFilePath?: string;
}

export const PrevStepButton = ({
  prevStep,
  openFilePath,
}: PrevStepButtonProps): JSX.Element => {
  const queryString = openFilePath
    ? `step=${prevStep}&openFilePath=${encodeURIComponent(openFilePath)}`
    : `step=${prevStep}`;

  return (
    <button type="button">
      <Link href={`./?${queryString}`}>prev step</Link>
    </button>
  );
};
