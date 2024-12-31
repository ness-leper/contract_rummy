import NextStep from "../components/NextStep";
import { PrintRules } from "../components/PrintRules";

export default function Step5() {
  return (
    <>
      <main className="flex min-h-screen flex-col items-center bg-gradient-to-b from-[#2e026d] to-[#15162c] text-white">
        <div className="mt-10">
          <h1 className="text-4xl">Round 5</h1>
          <div className="mt-10 text-xl">
            <PrintRules label="Deal 12 Cards" />
            <PrintRules label="Two Sets and 1 Sequence" />
          </div>
          <div className="mt-10">
            <NextStep label="Step 6" link="/step6" classes="" />
          </div>
        </div>
      </main>
    </>
  );
}
