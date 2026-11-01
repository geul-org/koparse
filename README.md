# koparse

Korean text analyzer for [GEUL](https://geul.org). It starts small — splitting
nouns from the particles (조사) and endings (어미) attached to them — and may grow
into a full morphological analyzer.

한국어 분석기. 어절에서 명사와 조사·어미를 떼어 내는 일부터 시작해, 필요하면 형태소 분석기로 키운다.

## 원칙

- 형태소 분석기·사전을 가져다 쓰지 않고 직접 만든다. 기존 분석기는 정확도 비교용으로만 쓴다.
- 사전은 CC0 원천(위키데이터 라벨·Lexeme)과 자체 생성 자료만 쓴다.
- 애매하면 정하지 않고 후보로 남긴다.

## 상태

2026-10-31: M0·M1·M2·M4 구현(fixture 어댑터). 설계는 [design.md](design.md).

- M0: 조사·어미 목록(`pkg/lexicon/data/suffixes.tsv`, 받침 조건·사슬 자리·무게), 지식 어댑터 인터페이스(`pkg/knowledge`), fixture 어댑터(`pkg/fixture`, `testdata/fixture.json`).
- M1: 정규화(NFC, 괄호 곁가지, 따옴표), 어절 분할, 분리 후보, 분리 판정(toulmin TAN).
- M2: 명사 단위(여러 어절 구간, 공백 변형, 복합어 조각), 명사 판정(TAN, R1~R4b + 분류 불일치 반박, 성질·관용 예외).
- M4: 문맥 통계(슬롯 묶음, 부호 있는 거리 구간, 같은 어절·절 경계 플래그).
- 아직 없음: SILK 어댑터(S1·S2·M3), 유일 후보 신뢰표(M5), GEUL 문서 출력(M6).

백과 표본 10건(fixture) 측정: 분리 정확도 96.5%, 명사 정밀도 94.8%(목표 95% 미달), 재현율 42.0%. 자세한 수치는 design.md §11.1.

## 사용

```sh
go build ./cmd/koparse
echo "고양이는 고기를 먹는다." | ./koparse -fixture testdata/fixture.json -topic Q146
```

문장마다 JSON 한 줄(design.md §5). `-config`로 Qualifier·공격 세기·문턱을 바꾸고, `-trace`로 어절 분리 후보 순위를 붙인다.

평가: `go run ./cmd/koparse-eval -fixture testdata/fixture.json -samples <표본.json> -gold-split <분리 정답.tsv> -gold-noun <명사 정답.tsv> [-judgments <손 판정.tsv>] [-out <디렉터리>]`. 표본 원문과 정답은 저장소 밖에 둔다.

## License

CC0 1.0 Universal. See [LICENSE](LICENSE) and [NOTICE](NOTICE).
