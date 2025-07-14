'use client'

import React, { useState } from "react";

function faktorial(n: number): number {
  if (n <= 1) return 1;
  return n * faktorial(n - 1);
}

function hitung(n: number): number {
  const numerator = faktorial(n);
  const denominator = Math.pow(2, n);
  return Math.ceil(numerator / denominator);
}

const HitungFaktorial: React.FC = () => {
  const [input, setInput] = useState<number>(0);
  const [hasil, setHasil] = useState<number | null>(null);

  const handleSubmit = (e: React.FormEvent) => {
    e.preventDefault();
    setHasil(hitung(input));
  };

  return (
    <div style={{ padding: "2rem", fontFamily: "sans-serif" }}>
      <h1>Hitung Faktorial Dibagi 2ⁿ (Ceil)</h1>
      <form onSubmit={handleSubmit}>
        <input
          type="number"
          value={input}
          onChange={(e) => setInput(parseInt(e.target.value))}
          min={0}
          placeholder="Masukkan nilai n"
        />
        <button type="submit" style={{ marginLeft: "0.5rem" }}>
          Hitung
        </button>
      </form>
      {hasil !== null && (
        <p>
          Hasil dari ⌈{input}! / 2^{input}⌉ adalah <strong>{hasil}</strong>
        </p>
      )}
    </div>
  );
};

export default HitungFaktorial;
