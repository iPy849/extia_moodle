document.querySelectorAll("tbody tr").forEach((tr) => {
  const domain = tr.querySelector("td:nth-child(1)").innerHTML;
  const data = tr.querySelector("td:nth-child(2)").innerHTML;

  // Clipboard copy
  tr.querySelector("td:nth-child(3) > button").addEventListener("click", () => {
    navigator.clipboard.writeText(data);
    alert(`La API-Key de ${domain.trim()} se ha copia correctamente`);
  });
});
