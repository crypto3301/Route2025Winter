#include <iostream>
#include <unordered_map>

int main()
{
    int n;
    std::cin >> n;
    std::vector<std::string> words(n);
    
    std::unordered_map<std::string, int> full_eqvivalent_words;
    for (auto& word : words)
    {
        full_eqvivalent_words[word]++;
    }
    
    std::unordered_map<std::string, int> uneven_eqvivalent_words;
    std::unordered_map<std::string, int> even_eqvivalent_words;
    for (auto& word : words)
    {
        /// fill even and uneven position of words
        std::string unevenWord;
        std::string evenWord;
        for (int i = 0; i < word.size(); ++i)
        {
            if (i % 2 == 1)
                unevenWords += word[i]
            else
                evenWord += word[i]
        }
        /// 
        
        
        /// check
        uneven_eqvivalent_words[unevenWord]++;
        even_eqvivalent_words[evenWord]++;
    }
    
    int sum = 0;
    for (auto it = uneven_eqvivalent_words.begin(); it != uneven_eqvivalent_words.end(); ++it)
    {
        sum += it->value;
    }
    for (auto it = even_eqvivalent_words.begin(); it != even_eqvivalent_words.end(); ++it)
    {
        sum += it->value;
    }
    for (auto it = full_eqvivalent_words.begin(); it != full_eqvivalent_words.end(); ++it)
    {
        sum -= it->value;
    }
    
    std::cout << sum << std::endl;
    return 0;
}
